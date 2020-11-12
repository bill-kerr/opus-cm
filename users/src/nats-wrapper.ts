import nats, { Stan } from 'node-nats-streaming';

class NatsWrapper {
  private _client?: Stan;

  get client() {
    if (!this._client) {
      throw new Error('Cannot access NATS client before connecting.');
    }
    return this._client;
  }

  private onConnect() {
    console.log('Connected to NATS.');
    this.client.on('close', () => this.onClose());
    process.on('SIGINT', () => this.client.close());
    process.on('SIGTERM', () => this.client.close());
  }

  private onClose() {
    console.log('NATS connection closed.');
    process.exit();
  }

  connect(clusterId: string, clientId: string, url: string) {
    this._client = nats.connect(clusterId, clientId, { url });
    return new Promise((resolve, reject) => {
      this.client.on('connect', () => {
        this.onConnect();
        resolve();
      });
      this.client.on('error', error => reject(error));
    });
  }
}

export const natsWrapper = new NatsWrapper();
