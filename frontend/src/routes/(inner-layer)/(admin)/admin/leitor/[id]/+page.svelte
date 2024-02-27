<script>
	import Card from '$lib/components/card.svelte';
	import Bold from '$lib/components/bold.svelte';

	export let data;

	let readerData = data.readerData,
		error = data.error;

	console.log(readerData);
</script>

<Card class="w-75">
	<h1 class="text-center">Dados de Usuário</h1>
	{#if readerData}
		<ul>
			<li>
				<Bold>Nome:</Bold>
				{readerData.User.name}
				{readerData.User.surname}
			</li>
			<li>
				<Bold>Email:</Bold>
				<a href={`mailto:${readerData.User.email}`}>{readerData.User.email}</a>
			</li>
			<li>
				<Bold>Data de nascimento:</Bold>
				{new Date(readerData.birthday).toLocaleDateString('pt-br')}
			</li>
			<li>
				<Bold>Endereço:</Bold>
				{readerData.address}
			</li>
			<li>
				<Bold>Número de telefone:</Bold>
				{readerData.phone_number}
			</li>
		</ul>
	{:else}
		<div class="text-danger">
			<p>Ops, aconteceu um erro. Talvez isso ajude:</p>
			<p>{error}</p>
		</div>
	{/if}
</Card>
