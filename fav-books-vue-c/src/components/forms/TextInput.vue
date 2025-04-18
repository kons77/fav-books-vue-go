<!-- converted to the Composition API -->
<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{label}}</label>
    <input
        :type="type"
        :name="name"
        :placeholder="placeholder"
        :required="required"
        :min="min"
        :max="max"
        :value="modelValue"
        :autocomplete="name + '-new'"
        @input="emit('update:modelValue', $event.target.value)"
        class="form-control">
      <div v-if="help" class="form-text">{{help}}</div>
  </div>
</template>



<script setup> 
  // export default{ name: 'TextInput', }

  const props = defineProps({
    name: String,
    type: String,
    label: String,
    placeholder: String,
    required: Boolean,
    min: Number,
    max: Number,
    modelValue: String,
    help: String,
  })

  const emit = defineEmits(['update:modelValue'])
</script>