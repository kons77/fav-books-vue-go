<!-- converted to the Composition API -->
<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{label}}</label>
    <select 
      :id="name" 
      class="form-select" 
      :name="name" 
      :required="required" 
      :multiple="multiple"
      :value="modelValue"
      @input="emit('update:modelValue', $event.target.value)">
      
      <option disabled value="">Choose...</option> <!-- no ability to choose on iPhone without this line -->
      <option v-for="option in items" :value="option.value" :key="option.value">
          {{option.text}}
      </option>
    </select>
  </div>
</template>


<script setup>
// export default { name: 'SelectInput', }
  const props = defineProps({
    items: {
      type: Array,
      required: true,
    }, 
    name : String, 
    required: Boolean, 
    label: String, 
    modelValue: [String, Array], // array for multiple select 
    multiple: Boolean, 
  })

  const emit = defineEmits(["update:modelValue"])
</script>