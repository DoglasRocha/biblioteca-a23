<script>
	import Card from '$lib/components/card.svelte';
</script>

<Card class="w-75">
	<h1 class="text-center">Administradores</h1>

	<div class="d-flex justify-content-center mt-5 px-5">

		<table class="table table-striped">
			<thead>
				<th>Nome</th>
			<tr>
				<td>
					<details>
						<summary>Admin1 Silva</summary>
	
					<p>Email : exemplo@email.com</p>
					<p>Status: Liberado</p>	
					<div>
						<button 
						type="button"
						class="btn btn-danger ms-1"
						>Revogar
						</button>
					</div> 
				</td>

			</tr>
			<tr>
				<td>
					<details>
						<summary>Admin2 Rocha</summary>
	
					<p>Email : exemplo2@email.com</p>
					<p>Status: Não Liberado</p>	
					<div>
						<button 
						type="button"
						class="btn btn-success ms-1"
						>Autorizar
						</button>
					</div>
				</td>
	
			</tr>
			</thead>
			<tbody>
								<!--

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
												//await api.delete(`/admin/livros/deletar/${book.id}`);
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
				-->
			</tbody>
		</table>
	</div>




		


	

</Card>

<style>
</style>
