<script>
</script>

<div class={'card ' + ($$restProps.class ?? '')}>
	<slot name="header" />
	<slot name="image" />
	<div class="card-body">
		<slot />
	</div>
</div>

<style>
</style>
