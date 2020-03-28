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
        v-for="item in items"
        :key="JSON.stringify(item)"
        flat
        bordered
        class="col-2 q-mb-md"
      >
        <q-card-section>
          <div class="row items-center">
            <div class="text-headline text-grey">{{ formatDate(item.created_at) }}</div>
            <q-space />
            <q-btn
              flat
              round
              color="primary"
              icon="fab fa-github-alt"
              type="a"
              target="__blank"
              :href="item.url"
            />
          </div>
          <div class="row items-end">
            <div class="text-h5 q-mt-sm q-mb-xs">{{ item.name }}</div>
            <q-space />
            <q-btn
              color="grey"
              round
              flat
              dense
              :icon="isExpanded(item.id) ? 'keyboard_arrow_up' : 'keyboard_arrow_down'"
              @click="toggleExpanded(item.id)"
            />
          </div>
          <div v-show="isExpanded(item.id)" class="q-pt-xs">
            <div class="text-body2 text-primary">{{ item.description }}</div>
          </div>
        </q-card-section>

        <q-separator v-if="item.tags && item.tags.length > 0" />

        <q-card-section v-if="item.tags && item.tags.length > 0">
          <q-chip
            v-for="tag in item.tags"
            :key="tag.id"
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
    </div>
  </q-page>
</template>

<script>
import { date } from 'quasar'

export default {
  data () {
    return {
      expandedId: '',
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
  mounted () {
    this.search()
  },
  methods: {
    isExpanded (id) {
      return this.expandedId === id
    },
    toggleExpanded (id) {
      if (this.isExpanded(id)) {
        this.expandedId = ''
      } else {
        this.expandedId = id
      }
    },
    formatDate (value) {
      const dt = new Date(value)
      return date.formatDate(dt, 'YYYY-MM-DD')
    },
    filterByTags (tags) {
      this.items = []
      this.tags = tags
      this.search()
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
        let data = []
        if (this.tags && this.tags.length > 0) {
          this.tags.forEach(async tagId => {
            const result = await this.$s.repo.searchByTag(tagId)
            const unique = []
            result.items.map(i => {
              const alreadyAdded = this.items.some(added => added.id === i.id)
              if (!alreadyAdded) {
                unique.push(i)
              }
            })
            this.items = this.items.concat(unique)
            this.pagination.rowsNumber = this.items.length
          })
        } else {
          data = await this.$s.repo.search({ pagination })
          this.items = this.items.concat(data.items)
        }
        const maxPages = Math.ceil(data.total / 10)
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
