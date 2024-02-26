<script>
	import Card from '$lib/components/card.svelte';
	import { api } from '$lib/utils/api.js';
	import { blur } from 'svelte/transition';
	export let data;

	let loan = data.loan,
		error = data.error;

	let returnDate = new Date(loan.return_date);
	let newReturnDate = new Date();
	let hasRenewed = false;

	newReturnDate.setDate(returnDate.getDate() + 7);

	const renewLoan = (loanId) => {
		try {
			let request = api.patch(`/renovar/${loanId}`);

			if (request.status == 200) {
				hasRenewed = true;
				loan = request.data;
			}
		} catch (err) {
			error = err.response.data;
		}
	};
</script>

<Card class="w-75">
	<h1 class="text-center">Renovar Empréstimo</h1>
	<div>
		{#if loan && !loan.has_returned}
			{#if !loan.has_renewed}
				<div class="mt-3" transition:blur>
					<p>
						Você pode renovar o livro <span class="fw-bold">{loan.Request.Book.name}</span> até o
						dia
						<span class="fw-bold">{newReturnDate.toLocaleDateString('pt-BR')}</span>.
					</p>
				</div>
				<div class="d-flex justify-content-end">
					<button
						class="btn btn-dark"
						on:click={() => {
							renewLoan(loan.id);
						}}>Renovar</button
					>
				</div>
			{:else if hasRenewed}
				<div transition:blur>
					<p>
						O livro <span class="fw-bold">{loan.Request.Book.name}</span> foi renovado com sucesso!
						Nova data de devolução:
						<span class="fw-bold">{new Date(loan.return_date).toLocaleDateString('pt-BR')}</span>
					</p>
				</div>
			{:else}
				<div transition:blur>
					<p>
						Você não pode renovar o livro <span class="fw-bold">{'João e o Pé de Feijão'}</span>,
						você deve devolvê-lo dia
						<span class="fw-bold">{new Date(loan.return_date).toLocaleDateString('pt-BR')}</span>.
					</p>
				</div>
			{/if}
		{:else}
			<p>Você não possui empréstimo em aberto, portanto, não tem o que renovar.</p>
		{/if}
		{#if error}
			<div class="text-danger">
				<p>Ops, ocorreu um erro. Talvez isso ajude:</p>
				<p>{error}</p>
			</div>
		{/if}
	</div>
</Card>
