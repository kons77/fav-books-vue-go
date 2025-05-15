<!-- converted to the Composition API -->
<template>
    <form 
      @submit.prevent="submit"
      ref="formRef"
      :name="name"
      :event="event"
      autocomplete="off" 
      :method="method" 
      :action="action" 
      class="needs-validation" 
      novalidate>
      <slot></slot>
    </form>
</template>


<script setup>
  import { ref } from 'vue';

  // in Vue 3 with <script setup>, the component name is automatically inferred from the file name
  // the name field can be omitted, or optionally declared using:
  //  script  export default { name: 'FormTag' } /script 


  const props = defineProps({
    method: String, 
    action: String,  
    name: String,  
    event: String,
  })
  
  const emit = defineEmits(["myevent", "submit", "validate"]);
  const formRef = ref(null);


  const submit = () => {
    if (formRef.value.checkValidity()) {
      emit(props.event); // emitting an event here instead of submitting the form
    }
    formRef.value.classList.add('was-validated');
  }

</script>