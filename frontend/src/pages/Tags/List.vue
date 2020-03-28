<template>
  <q-page class="q-pa-lg">
    <div class="text-h4 text-grey-8 q-mb-lg">
      {{ pagination.rowsNumber === 1 ? pagination.rowsNumber + ' tag' : pagination.rowsNumber + ' tags' }}
    </div>

    <q-table
      @request="search"
      :columns="columns"
      :data="items"
      :rows-per-page-options="[10, 30, 50, 0]"
      :pagination.sync="pagination"
      :loading="loading"
      square
      bordered
      flat
      row-key="id"
    >
      <template v-slot:top-right>
        <q-btn
          color="primary"
          icon="refresh"
          flat
          dense
          @click="search"
        />
      </template>
      <template slot="body" slot-scope="props" :props="props">
        <q-tr :props="props" @dblclick="edit(props.row)">
          <q-td key="name" :props="props">{{ props.row.name }}</q-td>
          <q-td key="total" :props="props">{{ props.row.total_repos || 0 }}</q-td>
          <q-td key="actions" :props="props">
            <q-btn
              color="primary"
              icon="edit"
              dense
              flat
              @click="edit(props.row)"
            />
            <q-btn
              color="negative"
              icon="close"
              dense
              flat
              @click="remove(props.row)"
            />
          </q-td>
        </q-tr>
      </template>
    </q-table>

    <q-dialog v-model="dialog">
      <q-card style="width: 500px">
        <q-card-section class="row items-center q-pb-none">
          <div class="text-h6">Edit tag</div>
          <q-space />
          <q-btn @click="close" icon="close" flat round dense v-close-popup />
        </q-card-section>
        <q-card-section>
          <q-input
            v-model="selected.name"
            filled
            autofocus
            @keyup.enter="updateTagName"
          />
        </q-card-section>
        <q-card-section class="row">
          <q-btn
            color="primary"
            label="Save"
            @click="updateTagName"
          />
          <q-space />
          <q-btn
            flat
            color="primary"
            label="Cancel"
            @click="close"
          />
        </q-card-section>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script>
export default {
  data () {
    return {
      columns: [
        { name: 'name', label: 'Name', align: 'left', field: 'name', sortable: true },
        { name: 'total', label: 'Repositories', align: 'left' },
        { name: 'actions', label: '', align: 'right', field: 'actions' }
      ],
      selected: {},
      dialog: false,
      items: [],
      loading: false,
      pagination: {
        page: 1,
        prevPage: 0,
        rowsPerPage: 10,
        rowsNumber: 10,
        sortBy: 'name'
      }
    }
  },
  mounted () {
    this.search()
  },
  methods: {
    edit (tag) {
      this.selected = { ...tag }
      this.dialog = true
    },
    close () {
      this.selected = {}
      this.dialog = false
    },
    async updateTagName () {
      try {
        const { selected } = this
        await this.$s.tag.update(selected)
        this.items = this.items.map(i => {
          if (i.id === selected.id) {
            return selected
          }
          return i
        })
        this.close()
      } catch (err) {
        this.$s.dialog.error(err)
      }
    },
    async remove (item) {
      try {
        await this.$s.tag.remove(item)
        this.items = this.items.filter(i => i.id !== item.id)
        this.pagination.rowsNumber -= 1
      } catch (err) {
        this.$s.dialog.error(err)
      }
    },
    async search (params) {
      this.loading = true
      try {
        const pagination = params && params.pagination ? params.pagination : this.pagination
        const filter = { pagination, total_repos: true }
        const data = await this.$s.tag.search(filter)
        this.items = data.items
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
