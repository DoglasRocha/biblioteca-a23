<script>
	import Card from '$lib/components/card.svelte';

	export let data;

	let book = data.book,
		error = data.error;
</script>

<Card class="w-75">
	{#if error}
		<h1>Erro!</h1>
		<p>Não há livro com este identificador</p>
	{:else}
		<h1 class="text-center">{book.name}</h1>
		<ul class="mt-3">
			<li>
				<span class="fw-bold">Descrição: </span>
				<p>
					{book.description}
				</p>
			</li>
			<li><span class="fw-bold">Gênero: </span>{book.gender}</li>
			<li><span class="fw-bold">Exemplares disponíveis: </span>{book.copies}</li>
		</ul>
		<div class="d-flex justify-content-end">
			<button class="btn btn-primary">Emprestar</button>
		</div>
	{/if}
</Card>
