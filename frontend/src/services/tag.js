export default class {
  constructor (deps) {
    this.$http = deps.http
  }

  async readAll () {
    const resp = await this.$http.get('/api/tag')
    return resp.data.data
  }

  async search (term) {
    const filter = {
      pagination: {
        per_page: 10,
        page: 1
      },
      ordering: {
        field: 'name',
        sort: 'asc'
      }
    }
    if (term) {
      filter.terms = [{
        id: '1',
        field: 'name',
        operator: 'contains',
        value: `%${term}%`
      }]
      filter.condition = '$1'
    }
    const resp = await this.$http.post('/api/tag/search', filter)
    return {
      items: resp.data.data || [],
      total: resp.data.meta.total || 0
    }
  }
}
