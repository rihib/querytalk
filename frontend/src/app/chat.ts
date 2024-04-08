export type MSG = {
  dbType: string;
  prompt: string;
};

export type VisualizableData = {
  chart: {
    type: string;
    x: string;
    y: string;
  };
  data: Array<{ [key: string]: string | number }>;
};

type ChatError = {
  code: number;
  message: string;
};

type ChatResponse = VisualizableData | ChatError;

export function isVisualizableData(response: any): response is VisualizableData {
  return response && typeof response.visualizableData === 'string';
}

export function isChatError(response: any): response is ChatError {
  return response && typeof response.code === 'number' && typeof response.message === 'string';
}

export async function chat(msg: MSG): Promise<VisualizableData | ChatError> {
  const response = await fetch('http://localhost:8080/v0.0.1/chat', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(msg as MSG)
  });

  if (!response.ok) {
    const message = `An error has occured: ${response.status}`;
    throw { code: response.status, message };
  }

  const data = await response.json();
  return data as ChatResponse;
}
