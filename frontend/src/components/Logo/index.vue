<template>
  <div>
    <img
      v-if="color === 'white'"
      src="~assets/logo-white.svg"
      :width="width || 250"
    />
    <img
      v-else
      src="~assets/logo-black.svg"
      :width="width || 250"
    />
  </div>
</template>

<script>
export default {
  props: ['color', 'width']
}
</script>
