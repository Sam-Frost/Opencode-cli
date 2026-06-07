# Brocode

A coding agent like Pi for your day-to-day AI coding tasks.  
Written in Go!

## Commands

### 1. Provider

- `provider list`  
  List all available providers.

- `provider login --p <provider_name> --apiKey <actual_api_key>`  
  Authenticate with a provider and save its API key.

### 2. Model

- `model list`  
  List all available models.

- `model -p <provider_name>`  
  List models for a specific provider.

- `model set <model_name>`  
  Set the active model.

### 3. Agent

- `agent -p "<user_query>"`  
  Execute the coding agent with the given prompt.
