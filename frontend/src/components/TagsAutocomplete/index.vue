<template>
  <div>
    <q-select
      v-model="selectedTags"
      filled
      use-input
      use-chips
      multiple
      hide-dropdown-icon
      input-debounce="0"
      option-value="id"
      option-label="name"
      new-value-mode="add-unique"
      @add="add"
      @new-value="createValue"
      @filter="search"
      :loading="loading"
      :label="label"
      :hint="hint"
      :options="allTags"
    >
      <template v-slot:append>
        <q-icon
          v-if="clear"
          name="cancel"
          @click.stop="clearAll"
          class="cursor-pointer q-mr-md"
        />
        <q-btn
          v-if="submit"
          flat
          icon="search"
          label="Search"
          dense
          color="primary"
          @click.stop="submit(selectedTags)"
        />
      </template>
    </q-select>
  </div>
</template>

<script>
export default {
  props: {
    label: {
      type: String,
      default: ''
    },
    hint: {
      type: String,
      default: ''
    },
    disable: {
      type: Boolean,
      default: false
    },
    tags: {
      type: Array
    },
    submit: {
      type: Function
    },
    clear: {
      type: Function
    }
  },
  data () {
    return {
      loading: false,
      allTags: [],
      selectedTags: []
    }
  },
  mounted () {
    if (this.tags && this.tags.length > 0) {
      this.selectedTags = this.tags
    }
  },
  methods: {
    clearAll () {
      this.selectedTags = []
      this.clear(this.selectedTags)
    },
    add (details) {
      this.selectedTags.push(details.value)
      this.$emit('added', details.value)
    },
    createValue (val, done) {
      if (val.length === 0) {
        return
      }
      const { selectedTags } = this
      val.split(/[,;|]+/)
        .map(v => v.trim())
        .filter(v => v.length > 0)
        .forEach(v => {
          if (!selectedTags.some(i => i.name === v)) {
            selectedTags.push({ name: v })
          }
        })

      done('')
      this.selectedTags = selectedTags
    },
    search (val, update, abort) {
      if (val.length < 1) {
        abort()
        return
      }
      update(async () => {
        this.loading = true
        try {
          const data = await this.$s.repo.searchTag(val)
          this.allTags = data.items || []
        } catch (err) {
          abort()
          console.error(err)
        }
        this.loading = false
      })
    }
  }
}
</script>
