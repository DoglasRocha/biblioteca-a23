<script>
	import Card from '$lib/components/card.svelte';
</script>

<Card class="w-75">
	<h1 class="text-center">Erro!</h1>
	<p>Algum erro aconteceu. Tente novamente mais tarde.</p>
	<p>
		Seu acesso ainda não foi liberado. Tente novamente mais tarde ou contate algum dos
		administradores.
	</p>
	<div class="d-flex justify-content-end">
		<a href="/login" class="btn btn-primary">Voltar para login</a>
	</div>
</Card>
