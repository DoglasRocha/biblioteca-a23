<script>
	export let name = '';
	export let type = 'text';
	export let label = '';
	export let placeholder = '';
	export let value = '';
	export let errorMsg = '';
	export let validation = /.*/;
	export let isInvalid = false;
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
  <FormField 
  	name="example" 
	type="email" 
	label="Email" 
	placeholder="example@email.com" 
	bind:value={value} 
	errorMsg="example"
	validation={/example/}
	bind:isInvalid={isInvalid}
	required
  />
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
		on:input={() => {
			isInvalid = !validation.test(value);
		}}
	/>
	{#if isInvalid}
		<p class="little text-danger">{errorMsg}</p>
	{/if}
</div>

<style>
	.little {
		font-size: 10px;
	}
</style>
