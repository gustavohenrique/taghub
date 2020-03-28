export default class {
  constructor (deps) {
    this.$http = deps.http
  }

  async searchByTag (id) {
    const resp = await this.$http.get(`/tags/${id}.json`)
    return {
      items: resp.data.data || [],
      total: resp.data.meta.total || 0
    }
  }

  async search (filter) {
    const { page } = filter.pagination
    const resp = await this.$http.get(`/repos/${page}.json`)
    return {
      items: resp.data.data || [],
      total: resp.data.meta.total || 0
    }
  }
}
