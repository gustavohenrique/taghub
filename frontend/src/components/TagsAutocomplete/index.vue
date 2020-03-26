<template>
  <q-select
    filled
    v-model="selected"
    use-input
    hide-selected
    fill-input
    hide-dropdown-icon
    input-debounce="0"
    option-value="id"
    option-label="name"
    :options="items"
    :loading="loading"
    :label="label"
    @filter="search"
    @input="setSelected"
  >
    <template v-slot:no-option>
      <q-item>
        <q-item-section class="text-grey">
          Não foi encontrado nenhum produto com esse nome ou SKU
        </q-item-section>
      </q-item>
    </template>
  </q-select>
</template>

<script>
export default {
  props: ['onSelect', 'label', 'horizontal', 'vertical'],
  data () {
    return {
      loading: false,
      selected: {},
      items: []
    }
  },
  computed: {
    service () {
      const services = {
        cast: this.$s.cast,
        questoes: this.$s.questoes,
        produto: this.$s.product
      }
      return services[this.horizontal]
    }
  },
  methods: {
    search (val, update, abort) {
      if (val.length < 1) {
        abort()
        return
      }
      update(async () => {
        try {
          this.loading = true
          const data = await this.service.search({
            term: val,
            vertical: this.vertical,
            sortBy: 'name'
          })
          this.items = data.items || []
        } catch (err) {
          abort()
          console.error(err)
        }
        this.loading = false
      })
    },
    setSelected (item) {
      this.$publish(this.onSelect, item)
      this.selected = {}
    }
  }
}
</script>
