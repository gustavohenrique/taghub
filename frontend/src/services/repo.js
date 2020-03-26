export default class {
  constructor (deps) {
    this.$http = deps.http
  }

  async getTotalStarredsRepositoriesToSync () {
    const resp = await this.$http.get('/api/repo/sync')
    return resp.data.data
  }

  async syncStarredRepositoriesNow () {
    const resp = await this.$http.post('/api/repo/sync')
    return resp.data.data
  }

  async search (params) {
    const filter = {
      pagination: {
        per_page: params.rowsPerPage,
        page: params.page
      },
      ordering: {
        field: params.sortBy,
        sort: params.descending ? 'desc' : 'asc'
      }
    }
    if (params.term) {
      filter.terms = [{
        id: '1',
        field: 'full_name',
        operator: 'icontains',
        value: `%${params.term}%`
      },
      {
        id: '2',
        field: 'email',
        operator: 'icontains',
        value: `%${params.term}%`
      },
      {
        id: '3',
        field: 'mobile',
        operator: 'contains',
        value: `%${params.term}%`
      },
      {
        id: '4',
        field: 'cpf',
        operator: 'contains',
        value: `%${params.term}%`
      }]
      filter.condition = '$1 or $2 or $3 or $4'
    }
    const resp = await this.$http.post('/student/search', filter)
    return {
      data: resp.data.data || [],
      total: resp.data.meta.total || 0
    }
  }

  async readOne (id) {
    const resp = await this.$http.get(`/student/${id}`)
    return resp.data.data
  }
}
