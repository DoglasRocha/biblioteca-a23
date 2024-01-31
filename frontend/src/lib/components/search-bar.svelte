<script>
	export let placeholder = '';
	export let value = '';
	export let onClick = () => {};
</script>

<!--
@component
This is a search bar component

Parameters:
- placeholder
- value
- onClick

  ```jsx
  <SearchBar placeholder="example@email.com" bind:value={value} onClick={() => console.log(value)}/>
	```
-->

<form class={'d-flex w-90 ' + ($$restProps.class ?? '')} role="search">
	<input class="form-control me-2" type="search" {placeholder} aria-label="Search" bind:value />
	<button class="btn btn-outline-success" type="submit" on:click={onClick}>Buscar</button>
</form>

<style>
	.w-90 {
		width: 90%;
	}
</style>
