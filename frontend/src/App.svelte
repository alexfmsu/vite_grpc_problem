<script>
	import { Button } from "@sveltestrap/sveltestrap";

	import google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb.js";

	import { LotsServiceClient } from "../proto/LotsService_grpc_web_pb";

	const clt = new LotsServiceClient("<host>:<port>");

	const req = new google_protobuf_empty_pb.Empty();
	const stream = clt.activeLots(req, {});

	let div = document.getElementById("div111");
	var elemDiv = document.createElement("div");

	elemDiv.style.cssText =
		"position:absolute;width:10%;height:10%;z-index:100;background:#FF0000;color:#000000;";
	document.body.appendChild(elemDiv);

	stream.on("data", function (response) {
		console.log(JSON.stringify(response));
		// alert('stream:data ' + JSON.stringify(response))

		let p = document.createElement("p");
		response = response.toObject();
		p.innerHTML = JSON.stringify(response["lot"]["stocksList"][0]["count"]);

		// {"lot":{"id":0,"desc":"Description","price":0}}
		// alert(JSON.stringify(response))

		// alert(JSON.stringify(response.toObject()))
		// Object.keys(response.array).forEach((prop)=> alert(prop));
		// alert(response)
		elemDiv.append(p);
		// elemDiv.append("Some text", p);

		// const msg = response.toObject()
		// AppendReceivedMsg(msg)
	});

	stream.on("status", function (status) {
		console.log(JSON.stringify(status));
		//alert('stream:status' + JSON.stringify(status))
	});

	stream.on("end", function (end) {
		console.log(JSON.stringify(end));
		// alert('stream:end' + JSON.stringify(end))
	});

	stream.on("error", function (err) {
		console.log(JSON.stringify(err));
		//alert('stream:error' + JSON.stringify(err))
	});
</script>

<svelte:head>
	<link
		rel="stylesheet"
		href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css"
	/>
</svelte:head>

<main>
	<h1>Hello {name}!</h1>
	<p>
		Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn
		how to build Svelte apps.
	</p>

	<Button color="primary" outline>Hello World!</Button>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
