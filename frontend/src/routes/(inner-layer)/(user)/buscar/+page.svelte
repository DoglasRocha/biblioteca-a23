<script>
	import Card from '$lib/components/card.svelte';
</script>

<Card class="w-75">
	<h1 class="text-center">Buscar livro</h1>
	<div class="d-flex justify-content-center mt-5">
		<form class="d-flex w-90" role="search">
			<input
				class="form-control me-2"
				type="search"
				placeholder="Pesquisar livro"
				aria-label="Search"
			/>
			<button class="btn btn-outline-success" type="submit">Buscar</button>
		</form>
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
								<li><span class="fw-bold"> Gênero: </span>aaaa</li>
								<li><span class="fw-bold">Exemplares disponíveis: </span>10</li>
								<li>
									<span class="fw-bold">Descrição: </span>
									<p>
										Lorem ipsum dolor sit, amet consectetur adipisicing elit. Laborum repellat
										recusandae praesentium quae magni quaerat necessitatibus commodi expedita
										maiores. Mollitia blanditiis numquam reprehenderit ex iste inventore, nobis
										repudiandae hic dolorum.
									</p>
								</li>
							</ul>
							<div class="d-flex justify-content-end">
								<a href="/emprestar/1" class="btn btn-dark">Emprestar</a>
							</div>
						</details>
					</td>
				</tr>
				<tr
					><td
						><details>
							<summary> Manifesto Comunista </summary>
							<ul>
								<li><span class="fw-bold">Gênero: </span>bbbb</li>
								<li><span class="fw-bold">Exemplares disponíveis: </span>5</li>
								<li>
									<span class="fw-bold">Descrição: </span>
									<p>
										Lorem ipsum dolor sit amet consectetur adipisicing elit. Ullam doloribus aut
										nesciunt, cum quia officia recusandae doloremque animi est numquam dicta esse
										blanditiis voluptatum rem accusamus totam ipsa dolorem labore!
									</p>
								</li>
							</ul>
							<div class="d-flex justify-content-end">
								<a href="/emprestar/2" class="btn btn-dark">Emprestar</a>
							</div>
						</details>
					</td></tr
				>
			</tbody>
		</table>
	</div>
</Card>

<style>
	.w-90 {
		width: 90%;
	}
</style>
