<script>
	export let name = '';
	export let type = 'text';
	export let label = '';
	export let placeholder = '';
	export let value = '';
	export let required = false;
	console.log(type);
</script>

<!--
@component
This is a generic form field component

Parameters:
- name
- type
- label
- placeholder
- value
  ```jsx
  <FormField name="example" type="email" label="Email" placeholder="example@email.com" bind:value={value} required/>
	```
-->

<div class="mt-3">
	<label for={name} class="form-label">{label}</label>
	<input
		{...$$restProps}
		{name}
		id={name}
		{...{ type }}
		{placeholder}
		bind:value
		class={'form-control ' + ($$restProps.class ?? '')}
		{required}
	/>
</div>
