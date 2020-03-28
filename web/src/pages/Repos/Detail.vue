<template>
  <q-page class="q-pa-lg">
    <div class="text-h4 text-grey-8 q-mb-lg">
      Editar Empresa
    </div>
    <q-form ref="form" class="q-pt-lg">
      <div class="field">
        <q-input
          input-class="cnpj"
          v-model="selected.cnpj"
          label="CNPJ"
          maxlength="14"
          hint="Somente números"
          filled
          :rules="[val => (val && val.length === 14) || 'CNPJ deve possuir 14 dígitos']"
        />
      </div>
      <div class="field q-pt-lg">
        <q-input
          input-class="social_name"
          v-model="selected.social_name"
          label="Razão Social"
          filled
        />
      </div>
    </q-form>
    <ActionsFooter
      :onOk="events.SAVED"
      :onCancel="events.BACK_TO_LIST"
    />
  </q-page>
</template>

<script>
import { events } from '../../constants'
export default {
  data () {
    return {
      events,
      selected: {}
    }
  },
  mounted () {
    const id = this.$route.params.id
    this.findById(id)

    const form = this.$refs.form
    this.$subscribe(events.SAVED, async () => {
      form.validate().then(async success => {
        if (success) {
          form.resetValidation()
          try {
            await this.$s.company.update(this.selected)
            this.$router.push({ name: 'companies' }).catch(() => {})
          } catch (err) {
            this.$s.dialog.error(err)
          }
        } else {
          this.$s.dialog.error('Preencha corretamente todos os campos.')
        }
      })
    })

    this.$subscribe(events.BACK_TO_LIST, () => {
      this.$router.replace({ name: 'companies' }).catch(() => {})
    })
  },
  destroyed () {
    this.$unsubscribe(events.SAVED)
    this.$unsubscribe(events.BACK_TO_LIST)
  },
  methods: {
    async findById (id) {
      try {
        this.selected = await this.$s.company.findById(id)
      } catch (err) {
        this.$s.dialog.error(err)
      }
    }
  }
}
</script>
