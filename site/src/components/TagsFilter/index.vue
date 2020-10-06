<template>
  <q-dialog
    v-model="visible"
    full-height
    full-width
  >
    <q-card
      flat
    >
      <q-bar class="bg-primary text-white">
        <span class="text-body2">Filter by {{ allTags.length }} tags</span>
        <q-space />
        <q-btn icon="close" flat round dense @click="visible = !visible" />
      </q-bar>
      <q-card-section class="q-pa-none q-ma-none">
        <q-tabs
          v-model="selectedTab"
          inline-label
          dense
        >
          <q-tab
            v-for="tab in tabs"
            :key="tab"
            :label="tab"
            :name="tab"
            class="bg-grey-3 q-pt-xs"
          >
            <q-badge style="display:none" color="primary" transparent floating>
              {{ totals[tab] }}
            </q-badge>
          </q-tab>
        </q-tabs>
        <q-tab-panels
          v-model="selectedTab"
        >
          <q-tab-panel
            v-for="tab in tabs"
            :key="tab"
            :name="tab"
          >
            <div v-if="getTagsByInitial(tab).length === 0">No tags found.</div>
            <div v-else class="text-body1 q-py-md">Select one or more tags and click on DONE.</div>
            <q-chip
              v-for="t in getTagsByInitial(tab)"
              text-color="white"
              clickable
              class="q-pa-md q-ma-sm"
              :key="t.id"
              :color="getColorAccordingOf(t)"
              @click="toggleSelectTag(t)"
            >
              {{ t.name }}
              <q-badge color="primary" transparent floating>
                {{ t.total_repos }}
              </q-badge>
            </q-chip>
          </q-tab-panel>
        </q-tab-panels>
      </q-card-section>
      <q-card-section class="text-center">
        <q-btn
          color="primary"
          label="Done"
          unelevated
          style="width:200px"
          @click="visible=false"
        />
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script>
import constants from '../../constants'

export default {
  data () {
    return {
      visible: false,
      totals: {},
      allTags: [],
      selectedIds: [],
      selectedTab: 'A',
      tabs: constants.INITIALS
    }
  },
  created () {
    this.$subscribe(constants.TAGS_FILTER_SHOW_DIALOG, () => {
      this.visible = !this.visible
    })
    this.readAll()
    this.$subscribe(constants.TAGS_FILTER_REFRESH, this.readAll)
  },
  destroyed () {
    this.$unsubscribe(constants.TAGS_FILTER_REFRESH)
    this.$unsubscribe(constants.TAGS_FILTER_SHOW_DIALOG)
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
      const totals = {}
      for (const initial of constants.INITIALS) {
        const tags = this.getTagsByInitial(initial)
        totals[initial] = tags.length
      }
      this.totals = totals
    },
    getTagsByInitial (initial) {
      const allTags = this.allTags
      if (initial === '#') {
        return allTags.filter(t => t.name.match(/^\d/))
      }
      return allTags.filter(t => t.name.toUpperCase().startsWith(initial))
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
