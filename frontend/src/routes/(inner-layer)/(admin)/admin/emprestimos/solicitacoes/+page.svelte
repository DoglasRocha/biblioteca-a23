<script>
	import Card from '$lib/components/card.svelte';
</script>

<Card class="w-75">
	<h1 class="text-center">Solicitações de empréstimo de livros</h1>

	<div class="d-flex justify-content-center mt-5 px-5">

		<table class="table table-striped">
			<thead>
				<tr>
					<th>Solicitações</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>
						<details>
							<summary>Livro 1 - Aluno 1</summary> 
							<div class="d-flex justify-content-end">
								<button 
								type="button"
								class="btn btn-success ms-1"
								>Liberar
								</button>
							</div> 
						</details>
					</td>
				</tr>
				<tr>
					<td>
						<details>
							<summary>Livro 1 - Aluno 1</summary> 
							<div class="d-flex justify-content-end">
								<button 
								type="button"
								class="btn btn-success ms-1"
								>Liberar
								</button>
							</div> 
						</details>
					</td>
				</tr>				
			</tbody>
		</table>
	</div>
</Card>
<style>
</style>
