<script>
	export let href;
</script>

<a {href} class={'fw-bold text-black ' + $$restProps.class ?? ''}><slot /></a>
