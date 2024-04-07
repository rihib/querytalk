export type MSG = {
  dbType: string;
  prompt: string;
};

type VisualizableData = {
  visualizableData: string;
};

type ChatError = {
  code: number;
  message: string;
};

type ChatResponse = VisualizableData | ChatError;

export async function chat(msg: MSG): Promise<ChatResponse> {
  const response = await fetch('http://localhost:8080/v0.0.1/chat', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(msg as MSG)
  });

  if (!response.ok) {
    const message = `An error has occured: ${response.status}`;
    throw { code: response.status, message } as ChatError;
  }

  const data = await response.json();
  return data as ChatResponse;
}
