#!/bin/bash

echo "Configurando Git hooks..."

# Ruta al directorio de githooks
HOOKS_DIR="./githooks"
GIT_HOOKS_DIR=".git/hooks"

# Verificar si estamos en un repositorio Git
if [ ! -d ".git" ]; then
  echo "Este script debe ejecutarse en la raíz de un repositorio Git."
  exit 1
fi

# Copiar los hooks
for hook in "$HOOKS_DIR"/*; do
  hook_name=$(basename "$hook")
  cp "$hook" "$GIT_HOOKS_DIR/$hook_name"
  chmod +x "$GIT_HOOKS_DIR/$hook_name"
  echo "Hook configurado: $hook_name"
done

echo "Git hooks configurados correctamente."
