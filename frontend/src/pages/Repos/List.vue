<template>
  <q-page class="q-pa-lg">
    <div class="text-h4 text-grey-8 q-mb-lg">
      Empresas
    </div>

    <div class="row col-12 justify-between q-mt-xl">
      <div class="column col">
        <SearchBox
          entity="Empresa"
          placeholder="Procure pelo CNPJ da empresa"
          :disable="true"
        />
      </div>
      <div class="flex row items-center">
        <q-btn
          round
          icon="add"
          class="inline"
          color="primary"
          size="large"
          @click="create"
        />
        <div class="text-body1 text-grey-8 text-weight-bold q-pl-lg q-mb-none">
          Criar Empresa
        </div>
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
      square
      bordered
      flat
      row-key="id"
      class="bg-grey-1 q-mt-lg"
      hide-bottom
    >
      <template v-slot:top-right="props">
        <q-btn
          color="primary"
          icon="refresh"
          flat
          dense
          @click="search"
        />
      </template>
      <template slot="body" slot-scope="props" :props="props">
        <q-tr
          :props="props"
        >
          <q-td key="cnpj" :props="props">{{ props.row.cnpj }}</q-td>
          <q-td key="social_name" :props="props">{{ props.row.social_name }}</q-td>
          <q-td key="actions" :props="props">
            <q-btn flat round color="primary" icon="edit" @click="edit(props.row.id)" />
          </q-td>
        </q-tr>
      </template>
    </q-table>
  </q-page>
</template>

<script>
import { date } from 'quasar'

export default {
  data () {
    return {
      columns: [
        { name: 'cnpj', label: 'CNPJ', align: 'left', field: 'cnpj' },
        { name: 'social_name', label: 'Razão Social', align: 'left', field: 'social_name' },
        { name: 'actions', label: '', align: 'right', field: 'actions' }
      ],
      items: [],
      filter: '',
      loading: false,
      pagination: {
        page: 1,
        prevPage: 0,
        rowsPerPage: 10,
        rowsNumber: 10
      }
    }
  },
  mounted () {
    this.search({
      pagination: this.pagination
    })
  },
  methods: {
    create () {
      this.$router.push({
        name: 'companies-create'
      })
    },
    edit (id) {
      this.$router.push({
        name: 'companies-detail',
        params: { id, edit: true }
      })
    },
    async remove (item) {
      try {
        await this.$s.company.remove(item)
        this.items = this.items.filter(i => i.id !== item.id)
      } catch (err) {
        this.$s.dialog.error(err)
      }
    },
    async search ({ pagination }) {
      this.loading = true
      try {
        pagination = pagination || this.pagination
        const data = await this.$s.company.search(pagination)
        this.items = data.items || []
        this.pagination = {
          ...pagination,
          rowsNumber: data.total
        }
      } catch (err) {
        this.$s.dialog.error(err)
      }
      this.loading = false
    },
    formatDate (data) {
      return date.formatDate(data, 'DD/MM/YYYY HH:mm')
    }
  }
}
</script>
