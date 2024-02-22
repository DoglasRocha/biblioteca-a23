<script>
	import Card from '$lib/components/card.svelte';
</script>

<Card class="w-75">
	<h1 class="text-center">Empréstimos ativos</h1>
	<div class="d-flex justify-content-center mt-5 px-5">
		<table class="table table-striped">
			<thead>
			</thead>
			<tbody>
				<tr>
                    <td>
                        <details>
                            <summary>Livro 1 - Aluno 1</summary>
                            <ul>
                                <li>Data de empréstimo: 31/12/1999</li>
                                <li>Data de devolução: 01/01/2000</li>
                            </ul>
                    </td>
                </tr>
                <tr>
                    <td>
                        <details>
                            <summary>Livro 2 - Aluno 2</summary>
                            <ul>
                                <li>Data de empréstimo: 31/12/1999</li>
                                <li>Data de devolução: 01/01/2000</li>
                            </ul>
                    </td>
                </tr>
			</tbody>
		</table>
	</div>
</Card>

<style>
</style>
