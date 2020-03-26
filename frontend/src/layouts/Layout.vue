<template>
  <q-layout class="bg-grey-1">
    <q-header elevated class="text-white" style="background: #344150" height-hint="61.59">
      <q-toolbar class="q-py-sm q-px-md">
        <div class="q-mr-xl q-pt-xs">
          <logo color="white" width="160" />
        </div>
        <q-select
          ref="search" dark dense standout use-input hide-selected
          class="GL__toolbar-select"
          color="black"
          :stack-label="false"
          label="Search by name or tag..."
          v-model="text"
          :options="filteredOptions"
          style="width: 300px"
        />

        <div class="q-ml-xs q-gutter-md text-body2 text-weight-bold row items-center no-wrap">
          <q-btn
            v-for="m in menus"
            :key="JSON.stringify(m)"
            :label="m.label"
            :to="m.route"
            flat
            text-color="white"
            no-caps
            type="a"
            class="text-md"
          />
        </div>

        <q-space />

        <div class="q-pl-sm q-gutter-sm row items-center no-wrap">
          <q-btn round dense flat :ripple="false" :icon="fabGithub" size="19px" color="white" no-caps />
        </div>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import { fabGithub } from '@quasar/extras/fontawesome-v5'

export default {
  data () {
    return {
      fabGithub,
      text: '',
      options: null,
      filteredOptions: []
    }
  },
  computed: {
    menus () {
      return [
        {
          label: 'Repositories',
          route: { name: 'sync' }
        },
        {
          label: 'Tags',
          route: { name: 'sync' }
        },
        {
          label: 'Synchronization',
          route: { name: 'sync' }
        }
      ]
    }
  }
}
</script>
