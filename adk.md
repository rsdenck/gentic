# https://adk.dev/get-started/go/
Create an agent project¶
Create an agent project with the following files and directory structure:


my_agent/
    agent.go    # main agent code
    .env        # API keys or project IDs
AI Models for ADK agents¶
root_agent = Agent(
    model=LiteLlm(model="ollama_chat/gemma3:latest"),
    name="dice_agent",
    description=(
        "hello world agent that can roll a dice of 8 sides and check prime"
        " numbers."
    ),
    instruction="""
      You roll dice and answer questions about the outcome of the dice rolls.
    """,
    tools=[
        roll_die,
        check_prime,
    ],
)export OLLAMA_API_BASE="http://localhost:11434"
adk web
