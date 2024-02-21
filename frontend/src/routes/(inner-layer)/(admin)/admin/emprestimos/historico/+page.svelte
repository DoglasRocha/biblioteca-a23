<script>
	import Card from '$lib/components/card.svelte';
</script>

<Card class="w-75">
	<h1 class="text-center">Histórico de empréstimo de livros</h1>

	<div class="d-flex justify-content-center mt-5 px-5">

		<table class="table table-striped">
			<thead>
				<th>Livros</th>
			<tr>
				<td>
					<details>
						<summary>Livro 1 - Aluno 1</summary>
				</td>
			</tr>
			<tr>				
                <td>
                <details>
                    <summary>Livro 2 - Aluno 2</summary>
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
