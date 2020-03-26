<template>
  <div class="row col-12 justify-between">
    <div class="column col-5">
      <p class="text-body1 text-grey-8">
        Pesquisa de {{ entity }}
      </p>
      <q-input
        id="searchWord"
        v-model="searchWord"
        class="col-5"
        filled
        :placeholder="placeholder"
        :disable="disable"
        @keyup.enter="$publish(evt, searchWord)"
        @keyup.esc="searchWord = ''"
      >
        <template v-slot:append>
          <q-icon name="search" />
        </template>
      </q-input>
    </div>
    <slot />
  </div>
</template>

<script>
export default {
  props: {
    entity: {
      type: String,
      default: 'itens'
    },
    placeholder: {
      type: String,
      default: 'Pesquise por nome'
    },
    disable: {
      type: Boolean,
      default: false
    },
    evt: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      searchWord: ''
    }
  }
}
</script>
