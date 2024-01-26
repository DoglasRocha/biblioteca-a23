<script lang="ts">
	export let name: string = '';
	export let type: string = 'text';
	export let label: string = '';
	export let placeholder: string = '';
	export let value: string = '';
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
