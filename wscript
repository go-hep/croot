#

top = '.'
out = '__build__'

def options(ctx):
    pass

def configure(ctx):
    ctx.load('go')
    
def build(ctx):

    ctx(
        features='cgopackage',
        name ='go-croot',
        source='''
        pkg/croot.go
        pkg/croot_genreflex.go
        pkg/croot_reflex.go
        pkg/croot_cintex.go
        ''',
        target='croot',
        use = [
            'croot',
            ],
        )

    ctx(
        features='go goprogram',
        name   = 'test-croot-ex-tree-00',
        source ='cmd/test-croot-ex-tree-00.go',
        target = 'test-croot-ex-tree-00',
        use = ['go-croot',],
        )
