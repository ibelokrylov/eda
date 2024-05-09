import { app } from './app';
import { config } from './config';

app.listen({ port: config.app.port, host: config.app.address }, (err) => {
	if (err) {
		process.exit(1)
	}
	console.log(`Server listen: ${config.app.address}:${config.app.port}`)
});
