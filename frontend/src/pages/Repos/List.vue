<template>
  <q-page class="q-pa-lg">
    <div class="text-h4 text-grey-8 q-mb-lg">
      {{ pagination.rowsNumber === 1 ? pagination.rowsNumber + ' repository' : pagination.rowsNumber + ' repositories' }}
    </div>

    <div class="q-pt-lg q-pb-lg">
      <tags-filter
        @updated="filterByTags"
      />
    </div>

    <div ref="cards">
      <q-card
        v-for="(item, index) in items"
        :key="JSON.stringify(item)"
        flat
        bordered
        class="col-2 q-mb-md"
      >
        <q-card-section>
          <div class="row items-center">
            <div>
              <div class="text-headline text-grey">{{ formatDate(item.created_at) }}</div>
              <div class="text-h5">{{ item.name }}</div>
            </div>
            <q-space />
            <q-btn
              flat
              dense
              color="primary"
              icon="more"
              @click="showDialogToAddTag(item, index)"
            />
            <q-btn
              flat
              dense
              color="grey"
              icon="open_in_new"
              type="a"
              target="__blank"
              class="q-ml-xs"
              :href="item.url"
            />
            <q-btn
              flat
              dense
              color="negative"
              icon="delete_forever"
              class="q-ml-lg"
              @click="remove(item)"
            />
          </div>
          <div class="q-pt-md">
            <div class="text-body2 text-primary">{{ item.description }}</div>
          </div>
        </q-card-section>

        <q-separator v-if="item.tags && item.tags.length > 0" />

        <q-card-section v-if="item.tags && item.tags.length > 0">
          <q-chip
            v-for="tag in item.tags"
            :key="tag.id"
            removable
            @remove="removeTagFromRepo(item, tag)"
            color="negative"
            text-color="white"
          >
            {{ tag.name }}
          </q-chip>
        </q-card-section>
      </q-card>

    </div>

    <div class="row justify-center q-my-md">
      <q-btn
        label="More"
        color="primary"
        @click="loadMore"
        :loading="loading"
        :disable="pagination.page >= pagination.maxPages"
      />
      <q-select
        filled
        label="Page"
        :options="pages"
        v-model="pagination.page"
        style="width:100px"
        class="q-ml-md"
      />

    </div>

    <tags-dialog
      :visible="dialog"
      :repo="selected"
      @close="closeDialogToAddTag"
      @added-tag="onAddedTag"
    />
  </q-page>
</template>

<script>
import { date } from 'quasar'
import constants from '../../constants'

export default {
  data () {
    return {
      expandedId: '',
      dialog: false,
      index: 0,
      selected: {},
      items: [],
      loading: false,
      tags: [],
      pagination: {
        page: 1,
        maxPages: 0,
        prevPage: 0,
        rowsPerPage: 10,
        rowsNumber: 10,
        sortBy: 'created_at',
        descending: true
      }
    }
  },
  async mounted () {
    await this.search()
  },
  computed: {
    pages () {
      const list = []
      for (let i = 1; i <= this.pagination.maxPages; i++) {
        list.push(i)
      }
      return list
    }
  },
  methods: {
    showDialogToAddTag (item, index) {
      const tags = item.tags ? [...item.tags] : []
      this.selected = {
        ...item,
        tags
      }
      this.index = index
      this.dialog = true
    },
    closeDialogToAddTag () {
      this.$publish(constants.TAGS_FILTER_REFRESH)
      this.dialog = false
      this.selected = {}
      this.index = 0
    },
    onAddedTag (tag) {
      const tags = this.items[this.index].tags || []
      tags.push(tag)
      this.items[this.index].tags = tags
    },
    formatDate (value) {
      const dt = new Date(value)
      return date.formatDate(dt, 'YYYY-MM-DD')
    },
    filterByTags (tags) {
      this.items = []
      this.tags = tags
      const pagination = {
        ...this.pagination,
        page: 1
      }
      this.search({ pagination })
    },
    async removeTagFromRepo (item, tag) {
      try {
        await this.$s.repo.removeTagFromRepo(item, tag)
        item.tags = item.tags.filter(i => i.id !== tag.id)
      } catch (err) {
        this.$s.dialog.error(err)
      }
    },
    async remove (item) {
      try {
        await this.$s.repo.remove(item)
        this.items = this.items.filter(i => i.id !== item.id)
      } catch (err) {
        this.$s.dialog.error(err)
      }
    },
    async loadMore () {
      const { pagination } = this
      if (pagination.page < pagination.maxPages) {
        pagination.page += 1
        await this.search({ pagination })
      }
    },
    async search (params) {
      this.loading = true
      try {
        const pagination = params && params.pagination ? params.pagination : this.pagination
        const filter = { pagination }
        let data = []
        if (this.tags && this.tags.length > 0) {
          filter.tags = this.tags
          data = await this.$s.repo.searchByTagsIds(filter)
        } else {
          data = await this.$s.repo.search(filter)
        }
        const maxPages = Math.ceil(data.total / pagination.rowsPerPage)
        this.items = this.items.concat(data.items) || []
        this.pagination = {
          ...pagination,
          maxPages,
          rowsNumber: data.total
        }
      } catch (err) {
        this.$s.dialog.error(err)
      }
      this.loading = false
    }
  }
}
</script>
