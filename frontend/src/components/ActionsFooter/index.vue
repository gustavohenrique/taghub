<template>
  <q-footer
    elevated
    class="row bg-white justify-between q-py-md q-px-xl"
  >
    <q-btn
      v-if="onCancel"
      id="btnCancel"
      color="primary"
      outline
      @click="$publish(onCancel)"
    >
      {{ labelCancel ? labelCancel : 'Cancelar' }}
    </q-btn>
    <q-btn
      v-if="onOk"
      id="btnOk"
      color="primary"
      @click="$publish(onOk)"
    >
      {{ labelOk ? labelOk : 'Salvar' }}
    </q-btn>
    <q-btn
      v-if="save"
      id="btnSave"
      color="primary"
      @click="save"
      label="Salvar"
      :disable="disable"
    />
  </q-footer>
</template>

<script>
export default {
  props: ['labelOk', 'labelCancel', 'onOk', 'onCancel', 'hideCancel', 'save', 'disable']
}
</script>
