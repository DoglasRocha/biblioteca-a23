<script>
	import Card from '$lib/components/card.svelte';
	import SearchBar from '$lib/components/search-bar.svelte';
	import { api } from '$lib/utils/api.js';
	export let data;

	let books = data.books,
		error = data.error;
	let searchBarValue;

	const searchBarOnClick = async () => {
		try {
			data = await api.get(`/livros/buscar?name=${searchBarValue}`);
			books = data.data;
			error = false;
		} catch (err) {
			error = true;
		}
	}
	let book = {
		name: 'Joao e o Pe de Feijao', description: "lorem",
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
						<tr>
							<td>
								<details>
									<summary>{book.name}</summary>
									<ul>
										<li><span class="fw-bold">Gênero: </span>{book.gender}</li>
										<li><span class="fw-bold">Exemplares disponíveis: </span>{book.copies}</li>
										<li>
											<span class="fw-bold">Descrição: </span>
											<p>{book.description}</p>
										</li>
									</ul>
									<div class="d-flex justify-content-end">
										<a href={`admin/livros/editar/1'${book.id}`} class="btn btn-secondary">Editar</a>
										<button type="button" class="btn btn-danger ms-1">Deletar</button>
									</div>
								</details>
							</td>
						</tr>
				{/if}
			</tbody>
		</table>
	</div>
</Card>

<style>
</style>

