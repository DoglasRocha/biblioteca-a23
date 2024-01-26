<script>
	export let name = '';
	export let type = 'text';
	export let label = '';
	export let placeholder = '';
	export let value = '';
</script>

<div class="mt-3">
	<label for={name} class="form-label">{label}</label>
	<input
		{type}
		{placeholder}
		on:input={(e) => {
			value = e.target.value;
		}}
		class="form-control"
	/>
</div>
