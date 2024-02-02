<script>
	import Card from '$lib/components/card.svelte';
	import SearchBar from '$lib/components/search-bar.svelte';
	import { fly } from 'svelte/transition';
	let searchedBook;
	let hasReadfromServer = 'Não';
	let hasRead = hasReadfromServer;
</script>

<Card class="w-75">
	<h1 class="text-center">Histórico de empréstimos</h1>
	<div class="d-flex justify-content-center mt-5">
		<SearchBar
			placeholder="Pesquisar livro"
			bind:value={searchedBook}
			onClick={() => console.log(searchedBook)}
		/>
	</div>
	<div class="w-90"></div>
	<div class="d-flex justify-content-center mt-5 px-5">
		<table class="table table-striped">
			<thead>
				<th>Nome</th>
			</thead>
			<tbody>
				<tr>
					<td>
						<details>
							<summary> João e o Pé de Feijão </summary>
							<ul>
								<li><span class="fw-bold">Gênero: </span>aaaa</li>
								<li>
									<span class="fw-bold">Descrição: </span>
									<p>
										Lorem ipsum dolor sit, amet consectetur adipisicing elit. Laborum repellat
										recusandae praesentium quae magni quaerat necessitatibus commodi expedita
										maiores. Mollitia blanditiis numquam reprehenderit ex iste inventore, nobis
										repudiandae hic dolorum.
									</p>
								</li>
								<li><span class="fw-bold">Data de empréstimo: </span>11/09/2001</li>
								<li><span class="fw-bold">Data de devolução: </span>18/09/2001</li>
								<li>
									<span class="fw-bold">Li o Livro:</span>

									<div>
										<input
											type="radio"
											name="has-read"
											id="not-read"
											value="Não"
											class="me-1"
											bind:group={hasRead}
										/><label for="has-read">Não</label>
										<input
											type="radio"
											name="has-read"
											id="has-read"
											value="Sim"
											class="ms-3 me-1"
											bind:group={hasRead}
										/><label for="has-read">Sim</label>
									</div>
								</li>
							</ul>
							{#if hasRead !== hasReadfromServer}
								<div
									transition:fly={{ delay: 0, duration: 250, x: -100, y: 0 }}
									class="d-flex justify-content-end"
								>
									<button class="btn btn-dark">Atualizar</button>
								</div>
							{/if}
						</details>
					</td>
				</tr>
			</tbody>
		</table>
	</div>
</Card>

<style>
	.w-90 {
		width: 90%;
	}
</style>
