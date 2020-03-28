<template>
  <q-page class="q-pa-lg">
    <div class="text-h4 text-grey-8 q-mb-lg">
      Criar Empresa
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
          :rules="[val => (val && val.trim().length === 14) || 'CNPJ deve possuir 14 dígitos']"
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
      :onOk="events.CREATED"
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
      selected: {
        cnpj: '',
        social_name: ''
      }
    }
  },
  mounted () {
    const form = this.$refs.form
    this.$subscribe(events.CREATED, () => {
      form.validate().then(async success => {
        if (success) {
          form.resetValidation()
          try {
            await this.$s.company.create(this.selected)
            this.$router.push({ name: 'companies' })
          } catch (err) {
            this.$s.dialog.error(err)
          }
        }
      })
    })

    this.$subscribe(events.BACK_TO_LIST, () => {
      this.$router.replace({ name: 'companies' }).catch(() => {})
    })
  },
  destroyed () {
    this.$unsubscribe(events.CREATED)
    this.$unsubscribe(events.BACK_TO_LIST)
  }
}
</script>
