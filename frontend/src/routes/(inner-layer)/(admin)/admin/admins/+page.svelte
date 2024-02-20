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
			</tbody>
		</table>
	</div>
</Card>


<style>
</style>
