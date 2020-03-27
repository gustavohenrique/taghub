<template>
  <q-page class="q-pa-lg">
    <div class="text-h4 text-grey-8 q-mb-lg">
      Repositories
    </div>

    <div class="row col-12 justify-between q-mt-xl">
      <div class="column col">
        <tags-autocomplete
          label="Search by tags"
          :submit="searchReposByTags"
          :clear="search"
        />
      </div>
    </div>

    <q-table
      @request="search"
      :columns="columns"
      :data="items"
      :dense="false"
      :filter="filter"
      :rows-per-page-options="[10, 30, 50, 0]"
      :pagination.sync="pagination"
      :loading="loading"
      bordered
      flat
      row-key="id"
      class="q-mt-lg"
    >
      <template slot="body" slot-scope="props" :props="props">
        <q-tr :props="props">
          <q-td key="name" :props="props">{{ props.row.name }}</q-td>
          <q-td key="description" :props="props">{{ props.row.description }}</q-td>
          <q-td key="tags" :props="props" align="right">
            <q-chip
              v-for="tag in props.row.tags"
              :key="tag.id"
              removable
              @remove="removeTagFromRepo(props.row, tag)"
              color="negative"
              text-color="white"
            >
              {{ tag.name }}
            </q-chip>
          </q-td>
          <q-td key="actions" :props="props" width="50">
            <q-btn
              flat
              round
              color="primary"
              icon="more"
              @click="showDialogToAddTag(props.row)"
            />
            <q-btn
              flat
              round
              color="primary"
              icon="fab fa-github-alt"
              type="a"
              target="__blank"
              :href="props.row.url"
            />
          </q-td>
        </q-tr>

        <tags-dialog
          :visible="dialog"
          :repo="selected"
          @close="dialog = false"
        />
      </template>
    </q-table>
  </q-page>
</template>

<script>
export default {
  data () {
    return {
      columns: [
        { name: 'name', label: 'Name', align: 'left', field: 'name', sortable: true },
        { name: 'description', label: 'Description', align: 'left' },
        { name: 'tags', label: 'Tags', align: 'right' },
        { name: 'actions', label: '', align: 'right' }
      ],
      dialog: false,
      selected: {},
      items: [],
      filter: '',
      loading: false,
      pagination: {
        page: 1,
        prevPage: 0,
        rowsPerPage: 10,
        rowsNumber: 10,
        sortBy: 'name',
        descending: false
      }
    }
  },
  mounted () {
    this.search({
      pagination: this.pagination
    })
  },
  methods: {
    showDialogToAddTag (item) {
      const tags = item.tags ? [...item.tags] : []
      this.selected = {
        ...item,
        tags
      }
      this.dialog = true
    },
    searchReposByTags (tags) {
      console.log('vou submter tags pra busscar:', tags)
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
    async search ({ pagination }) {
      this.loading = true
      try {
        pagination = pagination || this.pagination
        const data = await this.$s.repo.search(pagination)
        this.items = data.items || []
        this.pagination = {
          ...pagination,
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
