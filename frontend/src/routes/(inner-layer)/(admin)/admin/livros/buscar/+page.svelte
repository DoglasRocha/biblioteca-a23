<script>
	import Card from '$lib/components/card.svelte';
	import SearchBar from '$lib/components/search-bar.svelte';
	import { api } from '$lib/utils/api.js';
	import BookInTable from '$lib/components/book-in-table.svelte';
	export let data;

	let books = data.books,
		error = data.error;
	let searchBarValue;

	const searchBarOnClick = async () => {
		try {
			if (searchBarValue.trim() == '') data = await api.get(`/admin/livros/buscar`);
			else data = await api.get(`/admin/livros/buscar?name=${searchBarValue}`);

			books = data.data;
			error = false;
		} catch (err) {
			error = true;
		}
	};
</script>

<Card class="w-75">
	<h1 class="text-center">Buscar livro</h1>
	<div class="d-flex justify-content-center mt-5">
		<SearchBar
			placeholder="Pesquisar livro"
			bind:value={searchBarValue}
			onClick={searchBarOnClick}
		/>
	</div>
	<div class="d-flex justify-content-center mt-5 px-5">
		<table class="table table-striped">
			<thead>
				<th>Nome</th>
			</thead>
			<tbody>
				{#if error}
					<tr>
						<td> Não há livros!! </td>
					</tr>
				{:else}
					{#each books as book}
						<tr>
							<td>
								<BookInTable {book} admin={true}>
									<a href={`/admin/livros/editar/${book.id}`} class="btn btn-secondary"> Editar </a>
									<button
										type="button"
										class="btn btn-danger ms-1"
										on:click={async () => {
											let confirmation = confirm(
												'Você tem certeza que deseja deletar esse livro e todas as suas cópias?'
											);
											if (confirmation) {
												await api.delete(`/admin/livros/deletar/${book.id}`);
												document.location.href = document.location.href + '/';
											}
										}}>Deletar</button
									>
								</BookInTable>
							</td>
						</tr>
					{:else}
						<tr>
							<td> Não há livros!! </td>
						</tr>
					{/each}
				{/if}
			</tbody>
		</table>
	</div>
</Card>

<style>
</style>
