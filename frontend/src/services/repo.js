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
      id: tag.id,
      name: tag.name
    }
    const resp = await this.$http.post(`/api/repo/${repo.id}/tag`, req)
    return resp.data.data
  }

  async removeTagFromRepo (repo, tag) {
    await this.$http.delete(`/api/repo/${repo.id}/tag/${tag.id}`)
  }

  async searchByTagsIds (filter) {
    const { pagination, tags } = filter
    const req = {
      pagination: {
        per_page: pagination.rowsPerPage,
        page: pagination.page
      },
      ordering: {
        field: pagination.sortBy,
        sort: pagination.descending ? 'desc' : 'asc'
      },
      tags
    }
    const resp = await this.$http.post('/api/repo/tags/search', req)
    return {
      items: resp.data.data || [],
      total: resp.data.meta.total || 0
    }
  }

  async search (filter) {
    const { pagination } = filter
    const req = {
      pagination: {
        per_page: pagination.rowsPerPage,
        page: pagination.page
      },
      ordering: {
        field: pagination.sortBy,
        sort: pagination.descending ? 'desc' : 'asc'
      }
    }
    const resp = await this.$http.post('/api/repo/search', req)
    return {
      items: resp.data.data || [],
      total: resp.data.meta.total || 0
    }
  }
}
