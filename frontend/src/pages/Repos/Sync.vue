<template>
  <q-page class="q-pa-lg">
    <div class="text-h4 q-mb-lg text-custom-red">
      Synchronization
    </div>
    <div class="q-pt-lg text-body1">
      Fetch <b>{{ total }}</b> repositories that you starred and save them in the local database.
    </div>
    <div class="q-pt-lg">
      <q-btn
        color="primary"
        label="sync now"
        @click="syncNow"
      />
    </div>

    <div v-if="noResults" class="q-pt-lg text-body1">
      There are not new starred repositories to sync.
    </div>

    <div v-if="items && items.length > 0" class="q-pt-lg">
      <q-table
        title="Added repositories"
        :data="items"
        :columns="columns"
        bordered
        flat
        row-key="id"
        virtual-scroll
        :pagination.sync="pagination"
        :rows-per-page-options="[10, 0]"
      >
        <template slot="body" slot-scope="props" :props="props">
          <q-tr :props="props">
            <q-td key="name" :props="props">{{ props.row.name }}</q-td>
            <q-td key="url" :props="props">{{ props.row.url }}</q-td>
          </q-tr>
        </template>
      </q-table>
    </div>
  </q-page>
</template>

<script>
export default {
  data () {
    return {
      columns: [
        { name: 'name', label: 'Name', align: 'left' },
        { name: 'url', label: 'URL', align: 'left' }
      ],
      pagination: {
        rowsPerPage: 10
      },
      total: 0,
      noResults: false,
      items: []
    }
  },
  async created () {
    try {
      const total = await this.$s.repo.getTotalStarredsRepositoriesToSync()
      this.total = total
    } catch (err) {
      this.$s.dialog.error('I cannot get the total of starred repositories')
    }
  },
  methods: {
    async syncNow () {
      try {
        this.$q.loading.show()
        this.items = await this.$s.repo.syncStarredRepositoriesNow()
        this.noResults = true
      } catch (err) {
        let message = err.response.data.error
        if (err.response.status === 502) {
          message = 'A timeout ocurred on GitHub communication. Try again in few minutes.'
        }
        this.$s.dialog.error(message)
      }
      this.$q.loading.hide()
    }
  }
}
</script>
