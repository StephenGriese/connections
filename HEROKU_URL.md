# Finding Your Heroku App URL

Since you deploy from your personal computer, run these commands there:

## Option 1: Get the URL from Heroku CLI

```bash
cd ~/path/to/connections
heroku info
```

This will show your app details including the web URL.

## Option 2: Open in Browser Directly

```bash
heroku open
```

This will automatically open your app in your default browser.

## Option 3: List Your Apps

```bash
heroku apps
```

This shows all your Heroku apps. Your Connections app is likely named something like:
- `connections-solver-sjg`
- Or a random name like `quiet-mountain-12345`

Then the URL will be: `https://YOUR-APP-NAME.herokuapp.com`

## Option 4: Check Git Remote

```bash
cd ~/path/to/connections
git remote -v | grep heroku
```

Output will look like:
```
heroku  https://git.heroku.com/your-app-name.git (fetch)
heroku  https://git.heroku.com/your-app-name.git (push)
```

Your web URL is: `https://your-app-name.herokuapp.com`

## Option 5: Heroku Dashboard

Go to: https://dashboard.heroku.com/apps

You'll see all your apps listed. Click on your Connections app to see:
- App name
- URL (click "Open app" button in top right)

## Most Likely URL Format

Your app URL will be one of these formats:
- `https://connections-solver-sjg.herokuapp.com`
- `https://[random-name].herokuapp.com`

## Quick Command Reference

Run these on your **personal computer**:

```bash
# See all info including URL
heroku info

# Open app in browser
heroku open

# List all your apps
heroku apps

# Check app name from git
git remote -v
```

## Can't Find It?

If you haven't deployed yet:
1. The app might not exist yet
2. Check: `heroku apps` to see if you created one
3. If not, create one: `heroku create connections-solver-sjg`
4. Then deploy: `git push heroku main`

## After You Find It

Save it somewhere handy! Or just remember that `heroku open` will always open it for you.
