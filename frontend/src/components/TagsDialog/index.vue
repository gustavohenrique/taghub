<template>
  <q-dialog v-model="visible" persistent>
    <q-card style="width: 500px">
      <q-card-section class="row items-center q-pb-none">
        <div class="text-h6">{{ repo.name }}</div>
        <q-space />
        <q-btn @click="close" icon="close" flat round dense v-close-popup />
      </q-card-section>
      <q-card-section>
        <tags-autocomplete
          label="Add new tags"
          :tags="repo.tags"
          @added="addTagTo"
        />
      </q-card-section>
      <q-card-section class="row">
        <q-space />
        <q-btn
          flat
          color="primary"
          label="Discard"
          @click="close"
        />
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script>
export default {
  props: ['visible', 'repo'],
  data () {
    return {
    }
  },
  methods: {
    close () {
      this.$emit('close')
    },
    async addTagTo (tag) {
      try {
        const savedTag = await this.$s.repo.addTagToRepo(this.repo, tag)
        this.repo.tags.push(savedTag)
      } catch (err) {
        this.$s.dialog.error(err)
      }
    }
  }
}
</script>
