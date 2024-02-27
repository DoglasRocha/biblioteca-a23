<script>
	import Card from '$lib/components/card.svelte';
	import Bold from '$lib/components/bold.svelte';

	export let data;

	let readerData = data.readerData,
		error = data.error;
</script>

<Card class="w-75">
	<h1 class="text-center">Dados de Usuário</h1>
	{#if readerData}
		<ul>
			<li>
				<Bold>Nome:</Bold>
			</li>
			<li>
				<Bold>Email:</Bold>
			</li>
			<li>
				<Bold>Data de nascimento:</Bold>
			</li>
			<li>
				<Bold>Endereço:</Bold>
			</li>
			<li>
				<Bold>Número de telefone:</Bold>
			</li>
		</ul>
	{:else}
		<div class="text-danger">
			<p>Ops, aconteceu um erro. Talvez isso ajude:</p>
			<p>{error}</p>
		</div>
	{/if}
</Card>
