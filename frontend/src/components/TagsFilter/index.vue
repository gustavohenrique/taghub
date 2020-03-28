<template>
  <q-card
    flat
    style="background:#efefef"
  >
    <q-card-section>
      <q-btn
        flat
        dense
        color="grey-8"
        label="Filter by tags"
        :icon-right="show ? 'expand_less' : 'expand_more'"
        @click="show = !show"
      />
    </q-card-section>
    <q-card-section
      class="row q-gutter-md"
      v-if="show"
    >
      <div v-if="!allTags || allTags.length === 0">No tags found.</div>
      <q-chip
        v-for="t in allTags"
        text-color="white"
        clickable
        :key="t.id"
        :color="getColorAccordingOf(t)"
        @click="toggleSelectTag(t)"
      >
        {{ t.name }}
      </q-chip>
    </q-card-section>
  </q-card>
</template>

<script>
import constants from '../../constants'

export default {
  data () {
    return {
      show: false,
      allTags: [],
      selectedIds: []
    }
  },
  created () {
    this.readAll()
    this.$subscribe(constants.TAGS_FILTER_REFRESH, this.readAll)
  },
  destroyed () {
    this.$unsubscribe(constants.TAGS_FILTER_REFRESH)
  },
  watch: {
    selectedIds: {
      handler (val) {
        this.$emit('updated', val)
      },
      deep: true
    }
  },
  methods: {
    async readAll () {
      this.allTags = await this.$s.tag.readAll()
    },
    getColorAccordingOf (tag) {
      if (this.isSelected(tag.id)) {
        return 'negative'
      }
      return 'grey-6'
    },
    toggleSelectTag (tag) {
      if (this.isSelected(tag.id)) {
        this.selectedIds = this.selectedIds.filter(id => id !== tag.id)
        return
      }
      this.selectedIds.push(tag.id)
    },
    isSelected (id) {
      return this.selectedIds && this.selectedIds.some(i => i === id)
    }
  }
}
</script>
