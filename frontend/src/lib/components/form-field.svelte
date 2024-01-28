<script>
	export let name = '';
	export let type = 'text';
	export let label = '';
	export let placeholder = '';
	export let value = '';
	export let required = false;
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
		{name}
		id={name}
		{type}
		{placeholder}
		on:input={(e) => {
			value = e.target.value;
		}}
		class={'form-control ' + ($$restProps.class ?? '')}
		{required}
	/>
</div>
