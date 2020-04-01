<template>
  <q-layout view="hhh LpR fFf" class="bg-grey-1">
    <q-header
      elevated
      class="text-white"
      style="background: #344150"
      height-hint="61.59"
    >
      <q-toolbar class="q-py-sm q-px-md">
        <div class="q-pt-xs">
          <logo color="white" width="140" />
        </div>

        <q-space />

        <div class="q-pl-sm q-gutter-sm row items-center no-wrap">
          <q-btn
            round
            icon="fas fa-filter"
            flat
            dense
            no-caps
            @click="showTagsFilterDialog"
          />
          <q-btn
            round
            dense
            flat
            :icon="fabGithub"
            size="19px"
            color="white"
            no-caps
            type="a"
            target="__blank"
            href="https://github.com/gustavohenrique/taghub"
          />
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
import constants from '../constants'

export default {
  data () {
    return {
      fabGithub
    }
  },
  methods: {
    showTagsFilterDialog () {
      this.$publish(constants.TAGS_FILTER_SHOW_DIALOG)
    }
  }
}
</script>
