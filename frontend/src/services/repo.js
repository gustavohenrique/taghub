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

  async addTagToRepo (repo, tag) {
    const req = {
      id: repo.id,
      tag: {
        id: tag.id,
        name: tag.name
      }
    }
    const resp = await this.$http.post(`/api/repo/${repo.id}/tag`, req)
    return resp.data.data
  }

  async removeTagFromRepo (repo, tag) {
    await this.$http.delete(`/api/repo/${repo.id}/tag/${tag.id}`)
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
        field: 'name',
        operator: 'contains',
        value: `%${params.term}%`
      },
      {
        id: '2',
        field: 'description',
        operator: 'contains',
        value: `%${params.term}%`
      }]
      filter.condition = '$1 or $2'
    }
    const resp = await this.$http.post('/api/repo/search', filter)
    return {
      items: resp.data.data || [],
      total: resp.data.meta.total || 0
    }
  }

  async searchTag (term) {
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
    const resp = await this.$http.post('/api/repo/search/tag', filter)
    return {
      items: resp.data.data || [],
      total: resp.data.meta.total || 0
    }
  }
}
