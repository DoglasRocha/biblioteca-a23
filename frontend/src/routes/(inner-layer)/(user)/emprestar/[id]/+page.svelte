<script>
	import Card from '$lib/components/card.svelte';
	import { api } from '$lib/utils/api.js';

	export let data;

	let book = data.book,
		error = data.error;

	let errorFromServer = null;
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

		{#if book.copies >= 1}
			<div class="d-flex justify-content-end">
				<button
					class="btn btn-primary"
					on:click={async () => {
						try {
							let request = await api.post(`/emprestar/${book.id}`);

							if (request.status == 201) document.location.href = '/historico';
						} catch (error) {
							errorFromServer = error.response.data;
						}
					}}>Emprestar</button
				>
			</div>
		{/if}

		{#if errorFromServer}
			<div class="text-danger">
				<p>Ops. Ocorreu um erro. Talvez isso ajude:</p>
				<p>{errorFromServer}</p>
			</div>
		{/if}
	{/if}
</Card>
