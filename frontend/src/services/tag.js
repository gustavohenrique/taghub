export default class {
  constructor (deps) {
    this.$http = deps.http
  }

  async readAll () {
    const resp = await this.$http.get('/api/tag')
    return resp.data.data
  }

  async getTotalReposByTag (tag) {
    const resp = await this.$http.get(`/api/tag/${tag.id}/total`)
    return resp.data.data
  }

  async update (tag) {
    await this.$http.put(`/api/tag/${tag.id}`, tag)
  }

  async remove (tag) {
    await this.$http.delete(`/api/tag/${tag.id}`)
  }

  async search (filter) {
    const { term, pagination } = filter
    const req = {
      pagination: {
        per_page: pagination.rowsPerPage || 10,
        page: pagination.page || 1
      },
      ordering: {
        field: pagination.sortBy || 'name',
        sort: pagination.descending ? 'desc' : 'asc'
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
    const resp = await this.$http.post(`/api/tag/search?total_repos=${filter.total_repos}`, req)
    return {
      items: resp.data.data || [],
      total: resp.data.meta.total || 0
    }
  }
}
