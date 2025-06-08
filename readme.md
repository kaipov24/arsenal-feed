# Arsenal Feed

A Go web service that fetches Any Football Club data from the API-Football service(https://rapidapi.com/).

## Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/arsenal-feed.git
   cd arsenal-feed
   ```

2. **Create a `.env` file in the project root:**

   Copy the example and fill in your credentials:
   ```sh
   cp .env.example .env
   ```

   Edit `.env` and set your values:
   ```env
   API_FOOTBALL_KEY=your_api_football_api_key
   API_FOOTBALL_API_HOST=https://api-football-v1.p.rapidapi.com/v3
   API_FOOTBALL_API_TEAM=42
   ```

   - `API_FOOTBALL_KEY`: Your API-Football key from [RapidAPI](https://rapidapi.com/api-sports/api/api-football/).
   - `API_FOOTBALL_API_HOST`: The API-Football base URL.
   - `API_FOOTBALL_API_TEAM`: The team ID (e.g., Arsenal is `42`).

3. **Install dependencies:**
   ```sh
   go mod tidy
   ```

4. **Run the server:**
   ```sh
   go run main.go
   ```

## Endpoints

- `/next` — Get the next Arsenal fixture.
- `/last-five` — Get the last five Arsenal games.
- `/coach` — Get current coach info.
- `/team` — Get team information.
- `/standings` — Get current league standings.
- `/transfers` — Get latest transfers since January 1st, 2025.

