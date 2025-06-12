MIGRATE_VERSION="v4.18.2"
MIGRATE_PLATFORM="linux-amd64"

mkdir -p bin

echo "🔽 Downloading migrate ${MIGRATE_VERSION} for ${MIGRATE_PLATFORM}..."
curl -L "https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.${MIGRATE_PLATFORM}.tar.gz" | tar -xvz -C bin

chmod +x bin/migrate

echo "✅ Installed to: bin/migrate"
echo "🔍 Version:"
./bin/migrate -version

export PATH="./bin:$PATH"
echo "ℹ️  Temporary PATH set. You can run ./bin/migrate now."